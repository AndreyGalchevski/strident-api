import express, { Request, Response } from 'express';
import cors from 'cors';
import { json, urlencoded } from 'body-parser';
import { parse } from 'express-form-data';
import { config } from 'dotenv';

import connectToDB from './utils/db';

import authRouter from './auth/router';
import gigRouter from './gig/router';
import imageRouter from './image/router';
import lyricRouter from './lyric/router';
import memberRouter from './member/router';
import merchandiseRouter from './merchandise/router';
import songRouter from './song/router';
import videoRouter from './video/router';

const app = express();
const port = process.env.PORT || 8080;

config();
connectToDB();

app.use(cors({ origin: process.env.ALLOWED_ORIGIN }));

app.use(json({ limit: '1mb' }));
app.use(urlencoded({ extended: true, limit: '1mb' }));
app.use(parse());

app.get('/', (req: Request, res: Response) => {
  res.send('So Far, So Good... So What?');
});

app.use('/auth', authRouter);
app.use('/gigs', gigRouter);
app.use('/images', imageRouter);
app.use('/lyrics', lyricRouter);
app.use('/members', memberRouter);
app.use('/merchandises', merchandiseRouter);
app.use('/songs', songRouter);
app.use('/videos', videoRouter);

app.listen(port, () => console.log(`Server running on port ${port}`));
