import express from "express"
import Orde from "../modelos/orden.js"
import {auth } from "../midd/middleware.js"

const ruta = express.Router()

//Lista de todas las ordenes 
ruta.get("/", auth, async (req, res) => {
    const orders = await Orde.find();
    res.render("orders", {orders });
});

//Crea una orden 
ruta.post("/crear", async (req, res) => {

 const { productos, direccion } = req.body;

  if (!productos) return res.status(400).json({ ok: false, msg: "No hay productos" });

  await Orde.create({ productos, addres: direccion });

  res.status(200).json({ ok: true });
});

//Marca como entregada 
ruta.post("/:id/entregada", async (req, res) => {
    const id = req.params.id

    try{
        await Orde.findByIdAndUpdate(id, {estado: "entregada"})
        res.redirect("/panel")
    } catch (err){
        console.log(err)
        res.status(500).send("Error en la actualizacion")
    }
})

export default ruta