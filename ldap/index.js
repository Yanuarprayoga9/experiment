import express from "express"
import cors from "cors"
import dotenv from "dotenv"

dotenv.config()
const PORT = process.env.PORT || 3000;



const app = express()
app.use(express.json())

app.get("/wealth",(req,res)=>{
    res.send({msg:"app wealth "})
})

app.use(cors())


app.listen(PORT,(req,res)=>{
    console.log("app running on port ",PORT)
})