import { Component, OnInit } from '@angular/core';
import CardData from 'src/app/data/card.interface';
import SubjectEvent from 'src/app/data/subject-event.interface';
import Transformers from 'src/app/utils/transformers';
import { SubjectsService } from '../subjects.service';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';
import { SubjectEventsService } from '../subject-events.service';
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
  private subSubjectEvents: Subscription;
  private subDeleteSubjectEvent: Subscription;

  id: number;
  showAddEvent: boolean = false;

  constructor(private subjectsSvc: SubjectsService, private subjectEventsSvc: SubjectEventsService, private route: ActivatedRoute) {
    this.subjectEvents = []
  }

  subjectEventDescription: string = ""

  toggleShowAddEvent() {
    this.showAddEvent = !this.showAddEvent
  }

  onKey(event: any): void {
    this.subjectEventDescription = event.target.value
  }

  onSave(): void {
    let subjectEvent: SubjectEvent = {
      description: this.subjectEventDescription,
      subjectId: this.id,
      title: 'Title'
    }

    this.subjectEventsSvc.create(subjectEvent).subscribe(saved => {
      this.subjectEvents = [saved, ...this.subjectEvents]
    })

    this.subjectEventDescription = ""
    this.toggleShowAddEvent()
  }

  onDeleteSubjectEvent(subjectEvent: SubjectEvent){
    this.subDeleteSubjectEvent = this.subjectEventsSvc.delete(subjectEvent.ID).subscribe(() => {
      this.subjectEvents = this.subjectEvents.filter(sev => sev.ID !== subjectEvent.ID)
    })
  }

  ngOnInit() {
    this.subRouter = this.route.params.subscribe(params => {
      this.id = +params['id'];
      this.subSubjects = this.subjectsSvc.getSubjectById(this.id).subscribe(subject => {
        this.data = Transformers.transformSubjectToCardData(subject);
      });
      this.subSubjectEvents = this.subjectEventsSvc.getSubjectEventsBySubjectId(this.id).subscribe(subjecEvents => {
        this.subjectEvents = subjecEvents.reverse()
      })
    });
  }

  ngDestroy() {
    this.subRouter?.unsubscribe();
    this.subSubjects?.unsubscribe();
    this.subSubjectEvents?.unsubscribe();
    this.subDeleteSubjectEvent?.unsubscribe();
  }
}
