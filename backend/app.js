import express from "express"
import mongoose from "mongoose"
import cookieParser from "cookie-parser"
import dotenv from "dotenv"

import authrutas from "./rutas/usua_rutas.js"
import productoRutas from "./rutas/produc_rutas.js"
import ordenRutas from "./rutas/orde_rutas.js"

dotenv.config()

const app = express()

mongoose.connect(process.env.MONGO_URI)
    .then(() => console.log("MongoDB conectado"))

app.set("view engine", "pug")
app.set("views", "vistas")
app.use(express.urlencoded({ extended: true}))
app.use(cookieParser())
app.use(express.static("public"))

app.use("/auth", authrutas)
app.use("/productos", productoRutas)
app.use("/orden", ordenRutas)

app.get("/", (req,res)=>{
    res.redirect("/auth/login")
})

app.listen(3000, () => console.log("Servidor en http://localhost:3000"));