import express from "express"
import Orde from "../modelos/orden.js"
import {auth } from "../midd/middleware.js"

const ruta = express.Router()

ruta.get("/", auth, async (req, res) => {
    const orders = await Orde.find();
    res.render("orders", {orders });
});

ruta.post("/crear", async (req, res) => {

 const { productos, direccion } = req.body;

  if (!productos) return res.status(400).json({ ok: false, msg: "No hay productos" });

  await Orde.create({ productos, addres: direccion });

  res.status(200).json({ ok: true });
});


export default ruta