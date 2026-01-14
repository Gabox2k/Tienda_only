import jwt from 'jsonwebtoken'
import USER from '../modelos/usuario.js'

//Obtener el token y guarda en la cookie 
export const auth = async (req, res, next) =>{
    try{
        const token = req.cookies.token
        if (!token) return res.redirect("/auth/login")

       const verificaion = jwt.verify(token, process.env.JWT_SECRET)

       //Busca el usuario por su id 
       const usuario = await USER.findById(verificaion.id)

       if (!usuario) return res.redirect('/auth/login')

        req.usuario = usuario
        next() 
       } catch (error){
        console.log("No se obtuvo el usuario" )
        res.redirect("/auth/login")
    }
}
