import { ConfigService } from '@nestjs/config';

export const databaseConfig = (configService: ConfigService) => ({
  type: 'postgres',
  host: configService.get('DB_HOST'),
  port: +configService.get<number>('DB_PORT'),
  username: configService.get('DB_USER'),
  password: configService.get('DB_PASS'),
  database: configService.get('DB_NAME'),
  synchronize: true, // Set false in production
  entities: [__dirname + '/../**/*.entity.{ts,js}'],
});
