const express = require("express");
const mongo = require("mongoose");
const submit = require("./routes/submission.js");

const app = express();
app.listen(3000,() => console.log("The server is up and running"));

mongo.connect("mongodb://127.0.0.1:27017/cookoff_submission").then(() => console.log("Connected to DB"));

app.use("/",submit);

