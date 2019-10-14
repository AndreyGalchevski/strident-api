import { connect } from 'mongoose';

async function connectToDB(): Promise<void> {
  try {
    await connect(
      process.env.DB_URI,
      {
        useNewUrlParser: true,
        useFindAndModify: false,
        useUnifiedTopology: true,
      },
    );
    console.log('Connected to DB');
  } catch (error) {
    console.error(`Error connecting to DB: ${error}`);
  }
}

export default connectToDB;
