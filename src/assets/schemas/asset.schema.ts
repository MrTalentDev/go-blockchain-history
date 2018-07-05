import * as mongoose from 'mongoose';

export const AssetSchema = new mongoose.Schema({
    name: String,
    timestamp: Number,
    hash: String,
}, {versionKey: false});
