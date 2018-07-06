import * as mongoose from 'mongoose';
import {config} from 'dotenv';

config();

const local = 'mongodb://localhost:27017';
const env = process.env;
const prod = `mongodb://${env.LOGIN}:${env.PASSWORD}@${env.URL_CONNECTION}/${env.DB}`;

function getMongoProviderConfig(): string {
    return env.CONFIG === 'prod' ? prod : local;
}

export const databaseProviders = [
    {
        provide: 'DbConnectionToken',
        useFactory: async (): Promise<typeof mongoose> =>
            await mongoose.connect(getMongoProviderConfig()),
    },
];
