import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import Subject from '../data/subject.interface';

@Injectable({
  providedIn: 'root'
})
export class SubjectsService {
  apiURL: string = 'http://localhost:8080/subjects'

  constructor(private client: HttpClient) {
      
  }

  getSubjects() {
    return this.client.get<Subject[]>(`${this.apiURL}/`);
  }
  

  getSubjectById(id: number) {
    return this.client.get<Subject>(`${this.apiURL}/${id}`);
  }

  create(subject: Subject) {
    return this.client.post<Subject>(`${this.apiURL}/`, subject);
  }
}
