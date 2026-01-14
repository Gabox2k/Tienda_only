import mongoose from "mongoose"

//Guarda el orden en Mongo
const ordenSchema = new mongoose.Schema({
    productos: Array,
    addres: String,
    creacion: {type: Date, default: Date.now},
    estado: { type: String, default: "pendiente"}
})

export default mongoose.model("orden", ordenSchema)