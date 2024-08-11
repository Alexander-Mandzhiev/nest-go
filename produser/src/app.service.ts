import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';

@Injectable()
export class AppService {
  constructor(@Inject("SERVICE") private client: ClientProxy) { }

  async create() {
    const sanding = { hello: "hello" }
    return this.client.send({ cmd: "create_sand" }, sanding)
  }
}
