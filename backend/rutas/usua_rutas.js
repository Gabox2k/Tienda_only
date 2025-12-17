import express from "express"
import bcrypt from "bcryptjs"
import jwt from "jsonwebtoken"
import USER from "../modelos/usuario.js"

const ruta = express.Router()

ruta.get("/login" , (req, res) => {
    res.render("login")
})

ruta.post("/login", async (req, res) => {
    const { email, contra} = req.body
    const user = await USER.findOne({email}) 
    
    if (!user) return res.redirect("/auth/login")
    const storedHash = user.contra || user.contraseña || user.password
    if (!storedHash) return res.redirect("/auth/login")

    const contraHashe = await bcrypt.compare(contra, storedHash)
    if (!contraHashe) return res.redirect("/auth/login")
    
    const token = jwt.sign({ id: user._id}, process.env.JWT_SECRET)
    res.cookie("token", token, {httpOnly : true })
    console.log(`User logged in: ${user.email} role=${user.role}`)

    if (user.role === "admin"){
        console.log("Redirect a /panel")
        return res.redirect("/panel")
    } 
    return res.redirect("/productos")
})

ruta.get("/logout", (req,res) =>{
    res.clearCookie("token")
    res.redirect("/auth/login")
})

export default ruta