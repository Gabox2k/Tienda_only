import mongoose from "mongoose"

const ordenSchema = new mongoose.Schema({
    productos: Array,
    addres: String,
    creacion: {type: Date, default: Date.now},
    estado: { type: String, default: "pendiente"}
})

export default mongoose.model("orden", ordenSchema)