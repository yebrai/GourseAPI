package mooc

import (
	"GourseAPI/valueobject"
	"context"
	"errors"
)

// ErrEmptyCourseName is returned when the course name is empty.
var ErrEmptyCourseName = errors.New("the field Course Name cannot be empty")

// ErrEmptyDuration is returned when the course duration is empty.
var ErrEmptyDuration = errors.New("the field Duration cannot be empty")

// CourseName represents the name of a course.
type CourseName struct {
	value string
}

// NewCourseName creates a new CourseName instance
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{value: value}, nil
}

// String converts CourseName to a string.
func (name CourseName) String() string {
	return name.value
}

// CourseDuration represents the duration of a course.
type CourseDuration struct {
	value string
}

// NewCourseDuration creates a new CourseDuration instance,
func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	return CourseDuration{value: value}, nil
}

// String converts CourseDuration to a string.
func (duration CourseDuration) String() string {
	return duration.value
}

// CourseRepository defines the expected behavior of a course storage repository.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

// Course represents a course entity.
type Course struct {
	id       valueobject.UUID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new Course instance, validating its parameters.
func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := valueobject.NewUUID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

// ID returns the unique identifier of the course.
func (c Course) ID() valueobject.UUID {
	return c.id
}

// Name returns the name of the course.
func (c Course) Name() CourseName {
	return c.name
}

// Duration returns the duration of the course.
func (c Course) Duration() CourseDuration {
	return c.duration
}
