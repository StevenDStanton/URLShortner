'use strict';
import express, { Application } from 'express';
import dotenv from 'dotenv';

dotenv.config();

const app: Application = express();
const port = process.env.PORT;
if (!port) {
  //Done on purpose to blow up if I forget the .env file
  throw Error('Port is not defined');
}

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
