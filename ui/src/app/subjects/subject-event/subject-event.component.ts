import { Component, Input, Output, EventEmitter } from '@angular/core';
import SubjectEvent from 'src/app/data/subject-event.interface';

@Component({
  selector: 'app-subject-event',
  templateUrl: './subject-event.component.html',
  styleUrls: ['./subject-event.component.css']
})
export class SubjectEventComponent {
  @Input('data') data: SubjectEvent;

  @Output() 
  onDelete = new EventEmitter<number>();

  onDeleteAction() {
    this.onDelete.emit(this.data.ID)
  }

}
