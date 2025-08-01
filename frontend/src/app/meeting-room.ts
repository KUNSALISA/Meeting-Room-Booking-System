export interface RoomIn {
  ID: number;
  RoomName: string;
  Location: string;
  Capacity: number;
  Equipment: string[];
  TypeID: number;
  Types: RoomType;
  StatusID: number;
  Status: RoomStatus;
  Bookings: BookingIn[] | null;

  position: {
    x: number;
    y: number;
    width: number;
    height: number;
  };
}

export interface RoomType {
  ID: number;
  TypeName: string;
  Rooms: RoomIn[] | null;
}

export interface RoomStatus {
  ID: number;
  StatusName: string;
  Rooms: RoomIn[] | null;
}

export interface BookingIn{
  ID: number;
  Title: string;
  Description: string;
  Date: string;
  StartTime: string;
  EndTime: string;
  UserID: number;
  RoomID: number;
}

