import { Component, OnInit } from '@angular/core';
import CardData from 'src/app/data/card.interface';
import SubjectEvent from 'src/app/data/subject-event.interface';
import Transformers from 'src/app/utils/transformers';
import { SubjectsService } from '../subjects.service';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';
@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css']
})
export class ViewComponent implements OnInit {

  data: CardData = {} as CardData;
  subjectEvents: SubjectEvent[];
  private subRouter: Subscription;
  private subSubjects: Subscription;
  id: number;

  constructor(private subjectsSvc: SubjectsService, private route: ActivatedRoute) {
    this.subjectEvents = [
      {
        id: 1,
        subjectId: 1,
        title: "Subject Event",
        description: "Description",
        createdAt: new Date()
      },
      {
        id: 2,
        subjectId: 1,
        title: "Subject Event 2",
        description: "Description",
        createdAt: new Date()
      },
    ]
  }

  ngOnInit() {

    this.subRouter = this.route.params.subscribe(params => {
      this.id = +params['id'];
      this.subSubjects = this.subjectsSvc.getSubjectById(this.id).subscribe(subject => {
        this.data = Transformers.transformSubjectToCardData(subject);
      });
    });

    
  }

  ngDestroy() {
    this.subRouter?.unsubscribe();
    this.subSubjects?.unsubscribe();
  }
}
