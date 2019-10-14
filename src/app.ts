import express, { Request, Response } from 'express';
import cors from 'cors';
import { json } from 'body-parser';
import { config } from 'dotenv';

import connectToDB from './utils/db';

import authRouter from './auth/router';
import gigRouter from './gig/router';
import lyricRouter from './lyric/router';
import memberRouter from './member/router';
import songRouter from './song/router';
import videoRouter from './video/router';

const app = express();
const port = process.env.PORT || 8080;

config();
connectToDB();

app.use(cors({ origin: process.env.ALLOWED_ORIGIN }));
app.use(json());

app.get('/', (req: Request, res: Response) => {
  res.send('So Far, So Good... So What?');
});

app.use('/api/auth', authRouter);
app.use('/api/gigs', gigRouter);
app.use('/api/lyrics', lyricRouter);
app.use('/api/members', memberRouter);
app.use('/api/songs', songRouter);
app.use('/api/videos', videoRouter);

app.listen(port, err => {
  if (err) {
    console.error(err);
    return;
  }
  console.log(`server is listening on ${port}`);
});
