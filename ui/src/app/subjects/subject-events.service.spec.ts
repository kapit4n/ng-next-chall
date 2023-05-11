import { TestBed } from '@angular/core/testing';

import { SubjectEventsService } from './subject-events.service';

describe('SubjectEventsService', () => {
  let service: SubjectEventsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SubjectEventsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
