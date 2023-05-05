import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CardCustomComponent } from './card-custom.component';

describe('CardCustomComponent', () => {
  let component: CardCustomComponent;
  let fixture: ComponentFixture<CardCustomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CardCustomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CardCustomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
