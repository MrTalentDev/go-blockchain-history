import * as mongoose from 'mongoose';
import {config} from 'dotenv';

config();

const local = 'mongodb://localhost:27017';
const prod = `mongodb://${(process.env.LOGIN)}:${(process.env.PASSWORD)}@${(process.env.URL_CONNECTION)}/${(process.env.DB)}`;

function getMongoProviderConfig(): string{
    if (process.env.CONFIG === 'prod') {
        console.log('using prod DB mongo');
        return prod;
    }
    else if (process.env.CONFIG === 'local') {
        return local;
    }

    else {
        console.log('Using localhost DB by Default, if you want to use prod, please add CONFIG=prod to npm start');
        return local;
    }
}

export const databaseProviders = [
    {
        provide: 'DbConnectionToken',
        useFactory: async (): Promise<typeof mongoose> =>
            await mongoose.connect(getMongoProviderConfig()),
    },
];

