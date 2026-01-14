import express from "express"
import prod from "../modelos/productos.js"
import {auth } from "../midd/middleware.js"

const ruta = express.Router()

//Lista de productos 
ruta.get("/", auth, async (req, res) =>{
    const producto = await prod.find()
    res.render("productos", {producto})
})


ruta.get("/nuevo", auth, (req,res) =>{
    res.render("producto-form")
})


//Crea un producto 
ruta.post("/nuevo", auth, async(req,res) =>{
    await prod.create(req.body)
    res.redirect("/productos")
})


//Obtiene el producto editado
ruta.get("/editar/:id", auth, async (req, res)=>{
    const producto = await prod.findById(req.params.id)
    res.render ("producto-editar" , {producto})
})


//Edita el producto por id 
ruta.post("/editar/:id", auth, async (req, res) =>{
    const {precio} = req.body

    await prod.findByIdAndUpdate(req.params.id, {precio})

    res.redirect("/panel")
    
})


//Elimina el producto por su id 
ruta.get("/delete/:id", auth, async (req,res)=>{
    await prod.findByIdAndDelete(req.params.id)
    res.redirect("/panel")
})

export default ruta