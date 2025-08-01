import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { first } from 'rxjs';
import { RoomIn } from '../meeting-room';
import { MeetingRoomService } from '../meeting-room.service';

@Component({
  selector: 'app-room-map',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './room-map.component.html',
  styleUrl: './room-map.component.css',
})
export class RoomMapComponent {
  dataRoom: RoomIn[] = [];
  selectedFloor = 1;
  selectedRoom: RoomIn | null = null;

  constructor(private roomService: MeetingRoomService) {}

  mgOnInit() {
    this.roomService.getAllRooms().subscribe({
      next: (rooms) => {
        this.dataRoom = rooms;
      },
    });
  }

  selectFloor(floor: number) {
    this.selectedFloor = floor;
    this.selectedRoom = null;
  }

  getFloorNumberFromLocation(location: string): number {
    const match = location.match(/\d+/);
    return match ? parseInt(match[0], 10) : 0;
  }
  
  get roomsOfSelectedFloor(): RoomIn[] {
    return this.dataRoom.filter(
      (room) =>
        this.getFloorNumberFromLocation(room.Location) === this.selectedFloor
    );
  }
}
