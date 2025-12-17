import express from "express"
import brypt from "bcryptjs"
import jwt from "jsonwebtoken"
import USER from "../modelos/usuario.js"

const ruta = express.Router()

ruta.get("/login" , (req, res) => {
    res.render("logn")
})

ruta.post("/login", async (req, res) => {
    const { email, contra} = req.body
    const user = await USER.findOne({email}) 
    
    if (!user) return res.redirect("/auth/login")

    const contraHashe = await bcrypt.compare(contra, user.contra)
    if (!contraHashe) return res.redirect("/auth/login")
    
    const token = jwt.sign({ if: user._id}, process.env.JWT_SECRET)
    res.cookie("token", token)
    res.redirect("/productos")
})

ruta.get("/logout", (req,res) =>{
    res.clearCookie("token")
    res.redirect("/auth/login")
})

export default ruta