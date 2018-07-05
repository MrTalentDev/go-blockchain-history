import { Connection } from 'mongoose';
import { AssetSchema } from './schemas/asset.schema';

export const assetProviders = [
  {
    provide: 'AssetModelToken',
    useFactory: (connection: Connection) => connection.model('Asset', AssetSchema),
    inject: ['DbConnectionToken'],
  },
];
