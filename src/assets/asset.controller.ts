import {Controller, Get, Post, Body, Param} from '@nestjs/common';
import {CreateAssetDto} from './dto/create-asset.dto';
import {AssetService} from './asset.service';
import {Asset} from './interfaces/asset.interface';

//import {env} from '../../.env';

@Controller()
export class AssetController {
    constructor(private readonly assetService: AssetService) {
    }

    @Post()
    async create(@Body() createCatDto: CreateAssetDto) {
        this.assetService.create(createCatDto);
    }

    @Get('db/all')
    async findAll(): Promise<Asset[]> {
        return this.assetService.findAll();
    }

    @Get('/')
    answer(): string {
        return 'Ok';
    }
}
