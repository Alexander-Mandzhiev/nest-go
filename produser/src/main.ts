import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { NestExpressApplication } from '@nestjs/platform-express';
import * as cookieParser from 'cookie-parser'
import { Logger } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create<NestExpressApplication>(AppModule);
  app.disable('x-powered-by', 'X-Powered-By');

  app.setGlobalPrefix('api')
  app.use(cookieParser())

  app.enableCors({
    origin: [
      'http://localhost:3000',
      'http://127.0.0.1:3000',
      'https://localhost:3000',
      'https://127.0.0.1:3000'
    ],
    credentials: true,
    exposedHeaders: `set-cookie`
  })

  await app.listen(process.env.API_PORT)
  Logger.log(`server starting ${process.env.API_PORT}`)
}
bootstrap();
