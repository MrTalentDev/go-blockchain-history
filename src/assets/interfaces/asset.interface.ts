import { Document } from 'mongoose';

export interface Asset extends Document {
  readonly name: string;
  readonly timestamp: number;
  readonly hash: string;
}
