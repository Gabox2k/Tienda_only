import mongoose from "mongoose"

const usuarioSchema = new mongoose.Schema({
    nombre: String,
    email: { type: String, required: true, unique: true },
    contra: { type: String, required: true },
    role: { type: String, default: "user" }
})

export default mongoose.model("Usuario", usuarioSchema)
