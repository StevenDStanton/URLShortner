'use strict';
import express, { Application } from 'express';
import dotenv from 'dotenv';
import { indexRouter } from './routes/index';
import { urlRouter } from './routes/urlRouter';
import { initializeDb } from './lib/sqliteDb';

dotenv.config();
initializeDb()
  .then(() => {
    console.log('Database initialized');
  })
  .catch((error) => {
    console.error('Error initializing database:', error);
  });

const app: Application = express();
const port = process.env.PORT;
if (port === undefined) {
  throw Error('PORT is not defined');
}

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use('/', urlRouter);

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
