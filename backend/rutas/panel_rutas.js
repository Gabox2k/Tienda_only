import express from "express"
import { auth } from "../midd/middleware.js"
import Producto from  "../modelos/productos.js"
import Orden from "../modelos/orden.js"

const ruta = express.Router()

ruta.get("/", auth, async (req, res) => {
    if (!req.usuario || req.usuario.role !== "admin") return res.redirect("/auth/login")

    const productos = await Producto.find()
    const ordenes = await Orden.find()

    res.render("panel", { usuario: req.usuario, productos, ordenes })
})

export default ruta