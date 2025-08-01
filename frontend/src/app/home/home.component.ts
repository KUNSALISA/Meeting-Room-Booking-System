import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { DatePipe } from '@angular/common';
import { RoomMapComponent } from '../room-map/room-map.component';
import { SchedulerCardComponent } from '../scheduler-card/scheduler-card.component';

@Component({
  selector: 'app-home',
  imports: [RouterModule, DatePipe, RoomMapComponent, SchedulerCardComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  currentDate = new Date();

}
