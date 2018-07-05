import { Module } from '@nestjs/common';
import { AssetModule } from './assets/asset.module';

@Module({
  imports: [AssetModule],
})
export class ApplicationModule {}
