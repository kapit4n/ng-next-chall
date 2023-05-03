import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { SubjectsService } from '../subjects.service'

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

  data = [];

  constructor(private subjectsSvc: SubjectsService) { }

  ngOnInit() {
    this.subjectsSvc.getSubjects().subscribe(subjects => {
      this.data = subjects;
    });
  }

}
