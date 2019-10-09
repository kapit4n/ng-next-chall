import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';


interface Subject {
  id: number;
  name: string;
  catId: number;
}

@Injectable({
  providedIn: 'root'
})
export class SubjectsService {
  apiURL: string = 'http://localhost:8080/subjects'

  constructor(private client: HttpClient) {

  }

  getSubjects() {
    return this.client.get<Subject[]>(this.apiURL);
  }
}
