import moongoose from "mongoose"

const usuarioSchema= new moongoose.Schema({
    email: String,
    contraseña: String
})

export default moongoose.model("usuario", usuarioSchema)
