import jwt from 'jsonwebtoken'

export const auth = (req, res, next) =>{
    const token = req.cookies.token
    if (!token) return res.redirect("/auth/login")
}

try{
    jwt.verify(token, process.env.JWT_SECRET)
    next()
} catch {
    res.redirect("auth/login")
}