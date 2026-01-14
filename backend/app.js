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

//Conexion a Mongo 
const server = async () => {
  try {
    //Conexion correcta 
    await mongoose.connect(process.env.MONGO_URI)
    console.log("MongoDB conectado")

    //Asegura el admin
    const hashedcontra = await bcrypt.hash("admin123", 10)

    await USER.updateOne(
      { email: "admin@gmail.com" },
      {
        $set: {
          nombre: "admin",
          email: "admin@gmail.com",
          contra: hashedcontra,
          role: "admin"
        }
      },
      { upsert: true }
    )

    console.log("Admin asegurado")

  } catch (error) {
    console.log("Error al iniciar servidor:", error)
  }
}


server()

app.set("view engine", "pug")
app.set("views", "vistas")
app.use(express.static("public"))
app.use(express.json())
app.use(express.urlencoded({ extended: true}))
app.use(cookieParser())
app.use(express.static("public"))

app.use("/auth", authrutas)
app.use("/productos", productoRutas)
app.use("/orden", ordenRutas)
app.use("/ordenes", ordenRutas)
app.use("/panel", panelrutas)

app.get("/", (req,res)=>{
    res.redirect("/auth/login")
})

app.listen(3000, () => console.log("Servidor en http://localhost:3000"));