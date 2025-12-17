import express from "express"
import Orde from "../modelos/orden.js"
import {auth } from "../midd/middleware.js"

const ruta = express.Router()

ruta.get("/", auth, async (req, res) => {
    const orders = await Orde.find();
    res.render("orders", {orders });
});

export default ruta