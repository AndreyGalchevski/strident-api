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

app.use('/auth', authRouter);
app.use('/gigs', gigRouter);
app.use('/lyrics', lyricRouter);
app.use('/members', memberRouter);
app.use('/songs', songRouter);
app.use('/videos', videoRouter);

app.listen(port, () => console.log(`Server running on port ${port}`));
