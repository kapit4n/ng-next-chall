import { Component, OnInit } from '@angular/core';
import SubjectEvent from 'src/app/data/subject-event.interface';
import { DashboardService } from '../dashboard.service';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  subjectEvents: SubjectEvent[]
  sEventsSubscription: Subscription

  constructor(private dashboardSvc: DashboardService, private route: ActivatedRoute) {
    this.subjectEvents = []
  }

  ngOnInit() {
    this.sEventsSubscription = this.dashboardSvc.getSubjectEvents().subscribe(events => {
      // just display the top 3
      this.subjectEvents = events.reverse().slice(0, 3)
    })
  }
}
