import mongoose from "mongoose"

const ordenSchema = new mongoose.Schema({
    productos: Array,
    addres: String,
    creacion: {type: Date, default: Date.now}
})

export default mongoose.model("orden", ordenSchema)