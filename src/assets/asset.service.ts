import { Model } from 'mongoose';
import { Injectable, Inject } from '@nestjs/common';
import { Asset } from './interfaces/asset.interface';
import { CreateAssetDto } from './dto/create-asset.dto';

@Injectable()
export class AssetService {
  constructor(@Inject('AssetModelToken') private readonly assetModel: Model<Asset>) {}

  async create(createAssetDto: CreateAssetDto): Promise<Asset> {
    const assetService = new this.assetModel(createAssetDto);
    return await assetService.save();
  }

  async findAll(): Promise<Asset[]> {
    return await this.assetModel.find().exec();
  }
}
