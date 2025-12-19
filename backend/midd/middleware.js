import jwt from 'jsonwebtoken'
import USER from '../modelos/usuario.js'

export const auth = async (req, res, next) =>{
    try{
        const token = req.cookies.token
        if (!token) return res.redirect("/auth/login")

       const verificaion = jwt.verify(token, process.env.JWT_SECRET)
       const usuario = await USER.findById(verificaion.id)

       if (!usuario) return res.redirect('/auth/login')

        req.usuario = usuario
        next() 
       } catch (error){
        console.log("error auth:" )
        res.redirect("/auth/login")
    }
}
