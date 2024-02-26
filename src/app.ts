'use strict';
import express, { Application } from 'express';
import dotenv from 'dotenv';
import { urlRouter } from './routes/urlRouter';

dotenv.config();

const app: Application = express();
const port = process.env.PORT;
if (!port) {
  //Done on purpose to blow up if I forget the .env file
  throw Error('Port is not defined');
}

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use('/api/url', urlRouter);

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
