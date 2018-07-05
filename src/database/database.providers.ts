import * as mongoose from 'mongoose';
import {config} from 'dotenv';
config();

export const databaseProviders = [
    {
        provide: 'DbConnectionToken',
        useFactory: async (): Promise<typeof mongoose> =>
            await mongoose.connect(`mongodb://${(process.env.LOGIN)}:${(process.env.PASSWORD)}@${(process.env.URL_CONNECTION)}/${(process.env.DB)}`),
    },
];
