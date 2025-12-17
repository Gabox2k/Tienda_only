import express from "express"
import mongoose from "mongoose"
import cookieParser from "cookie-parser"
import dotenv from "dotenv"

import authrutas from "./rutas/usua_rutas.js"
import productoRutas from "./rutas/produc_rutas.js"
import ordenRutas from "./rutas/orde_rutas.js"
import panelrutas from "./rutas/panel_rutas.js"
import USER from "./modelos/usuario.js"
import bcrypt from "bcryptjs/dist/bcrypt.js"

dotenv.config()

const app = express()

const server = async() => {
    try {
        mongoose.connect(process.env.MONGO_URI)
        .then(() => console.log("MongoDB conectado"))

        const admin = await USER.findOne({ email: "admin@gmail.com"})
        if (!admin){
            const hashedcontra = await bcrypt.hash("admin123", 10)
            await USER.create({
                nombre: "admin",
                email: "admin@gmail.com",
                contra: hashedcontra,
                role: "admin"
            })
            console.log("Admin creado")}

    } catch (error){
        console.log("error al iniciar servidor:", error)
    }
}

server()

app.set("view engine", "pug")
app.set("views", "vistas")
app.use(express.urlencoded({ extended: true}))
app.use(cookieParser())
app.use(express.static("public"))

app.use("/auth", authrutas)
app.use("/productos", productoRutas)
app.use("/orden", ordenRutas)
app.use("/panel", panelrutas)

app.get("/", (req,res)=>{
    res.redirect("/auth/login")
})

app.listen(3000, () => console.log("Servidor en http://localhost:3000"));