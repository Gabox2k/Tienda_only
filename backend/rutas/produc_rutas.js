import express from "express"
import prod from "../modelos/productos.js"
import {auth } from "../midd/middleware.js"

const ruta = express.Router()

ruta.get("/", auth, async (req, res) =>{
    const producto = await prod.find()
    res.render("productos", {producto})
})

ruta.get("/nuevo", auth, (req,res) =>{
    res.render("producto-form")
})

ruta.post("/nuevo", auth, async(req,res) =>{
    await prod.create(req.body)
    res.redirect("/producto")
})

ruta.get("/delete/:id", auth, async (req,res)=>{
    await prod.findIdAndDelete(req.params.id)
    res.redirect("/productos")
})

export default ruta