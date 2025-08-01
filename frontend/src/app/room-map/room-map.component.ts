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
  selectedRoomDetail: RoomIn | null = null;

  constructor(private roomService: MeetingRoomService) {}

  ngOnInit() {
    this.roomService.getAllRooms().subscribe({
      next: (rooms) => {
        console.log('Rooms fetched:', rooms);
        this.dataRoom = rooms;
      },
      error: (err) => {
        console.error('Failed to load rooms', err);
      },
    });
  }

  selectFloor(floor: number) {
    console.log('Selected floor:', floor);
    this.selectedFloor = floor;
    this.selectedRoomDetail = null;
    console.log('Rooms of new floor:', this.roomsOfSelectedFloor);
  }

  getFloorNumberFromLocation(location: string): number {
    const match = location.match(/\d+/);
    return match ? parseInt(match[0], 10) : 0;
  }

  get roomsOfSelectedFloor(): RoomIn[] {
    const filtered = this.dataRoom.filter(
      (room) =>
        this.getFloorNumberFromLocation(room.Location) === this.selectedFloor
    );
    console.log('Rooms of floor', this.selectedFloor, filtered);
    return filtered;
  }

  loadRoomDetail(roomId: number) {
    this.roomService.getRoomById(roomId.toString()).subscribe({
      next: (roomDetail) => {
        this.selectedRoomDetail = roomDetail;
      },
      error: (err) => {
        console.error('Error loading room detail', err);
        this.selectedRoomDetail = null;
      },
    });
  }
}
