import { ApiProperty } from '@nestjs/swagger';
import { IsString, IsNumber, IsEnum } from 'class-validator';

enum TaskStatus {
  TODO = 'TODO',
  IN_PROGRESS = 'IN_PROGRESS',
  DONE = 'DONE',
}

export class CreateTaskDto {
  @ApiProperty({ example: 'Buy groceries', description: 'Task title' })
  @IsString()
  title: string;

  @ApiProperty({
    example: 'Buy milk and bread',
    description: 'Task description',
  })
  @IsString()
  description: string;

  @ApiProperty({ example: 1, description: 'Task priority' })
  @IsNumber()
  priority: number;

  @ApiProperty({
    enum: TaskStatus,
    example: 'TODO',
    description: 'Task status',
  })
  @IsEnum(TaskStatus)
  status: string;
}
