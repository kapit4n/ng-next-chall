import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubjectEventComponent } from './subject-event.component';

describe('SubjectEventComponent', () => {
  let component: SubjectEventComponent;
  let fixture: ComponentFixture<SubjectEventComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubjectEventComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SubjectEventComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
