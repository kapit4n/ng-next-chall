import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import SubjectEvent from '../data/subject-event.interface';

@Injectable({
  providedIn: 'root'
})
export class SubjectEventsService {
  apiURL: string = 'http://localhost:8080/subject_events'

  constructor(private client: HttpClient) {
  }

  getSubjectEvents() {
    return this.client.get<SubjectEvent[]>(`${this.apiURL}/`);
  }

  getSubjectEventsBySubjectId(subjecId: number) {
    return this.client.get<SubjectEvent[]>(`${this.apiURL}/?subjectId=${subjecId}`);
  }

  getSubjectEventById(id: number) {
    return this.client.get<SubjectEvent>(`${this.apiURL}/${id}`);
  }

  create(subjectEvent: SubjectEvent) {
    return this.client.post<SubjectEvent>(`${this.apiURL}/`, subjectEvent);
  }

}
