import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import SubjectEvent from '../data/subject-event.interface';

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
  eventsURL: string = 'http://localhost:8080/subject_events'


  constructor(private client: HttpClient) {
  }

  getSubjectEvents() {
    return this.client.get<SubjectEvent[]>(`${this.eventsURL}/`);
  }

}
