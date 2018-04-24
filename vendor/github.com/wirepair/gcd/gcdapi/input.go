// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Input functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// No Description.
type InputTouchPoint struct {
	X             float64 `json:"x"`                       // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y             float64 `json:"y"`                       // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX       float64 `json:"radiusX,omitempty"`       // X radius of the touch area (default: 1.0).
	RadiusY       float64 `json:"radiusY,omitempty"`       // Y radius of the touch area (default: 1.0).
	RotationAngle float64 `json:"rotationAngle,omitempty"` // Rotation angle (default: 0.0).
	Force         float64 `json:"force,omitempty"`         // Force (default: 1.0).
	Id            float64 `json:"id,omitempty"`            // Identifier used to track touch sources between events, must be unique within an event.
}

type Input struct {
	target gcdmessage.ChromeTargeter
}

func NewInput(target gcdmessage.ChromeTargeter) *Input {
	c := &Input{target: target}
	return c
}

type InputDispatchKeyEventParams struct {
	// Type of the key event.
	TheType string `json:"type"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp float64 `json:"timestamp,omitempty"`
	// Text as generated by processing a virtual key code with a keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events (default: "")
	Text string `json:"text,omitempty"`
	// Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
	UnmodifiedText string `json:"unmodifiedText,omitempty"`
	// Unique key identifier (e.g., 'U+0041') (default: "").
	KeyIdentifier string `json:"keyIdentifier,omitempty"`
	// Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Code string `json:"code,omitempty"`
	// Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	Key string `json:"key,omitempty"`
	// Windows virtual key code (default: 0).
	WindowsVirtualKeyCode int `json:"windowsVirtualKeyCode,omitempty"`
	// Native virtual key code (default: 0).
	NativeVirtualKeyCode int `json:"nativeVirtualKeyCode,omitempty"`
	// Whether the event was generated from auto repeat (default: false).
	AutoRepeat bool `json:"autoRepeat,omitempty"`
	// Whether the event was generated from the keypad (default: false).
	IsKeypad bool `json:"isKeypad,omitempty"`
	// Whether the event was a system key event (default: false).
	IsSystemKey bool `json:"isSystemKey,omitempty"`
	// Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
	Location int `json:"location,omitempty"`
}

// DispatchKeyEventWithParams - Dispatches a key event to the page.
func (c *Input) DispatchKeyEventWithParams(v *InputDispatchKeyEventParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchKeyEvent", Params: v})
}

// DispatchKeyEvent - Dispatches a key event to the page.
// type - Type of the key event.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred.
// text - Text as generated by processing a virtual key code with a keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events (default: "")
// unmodifiedText - Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
// keyIdentifier - Unique key identifier (e.g., 'U+0041') (default: "").
// code - Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
// key - Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
// windowsVirtualKeyCode - Windows virtual key code (default: 0).
// nativeVirtualKeyCode - Native virtual key code (default: 0).
// autoRepeat - Whether the event was generated from auto repeat (default: false).
// isKeypad - Whether the event was generated from the keypad (default: false).
// isSystemKey - Whether the event was a system key event (default: false).
// location - Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
func (c *Input) DispatchKeyEvent(theType string, modifiers int, timestamp float64, text string, unmodifiedText string, keyIdentifier string, code string, key string, windowsVirtualKeyCode int, nativeVirtualKeyCode int, autoRepeat bool, isKeypad bool, isSystemKey bool, location int) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchKeyEventParams
	v.TheType = theType
	v.Modifiers = modifiers
	v.Timestamp = timestamp
	v.Text = text
	v.UnmodifiedText = unmodifiedText
	v.KeyIdentifier = keyIdentifier
	v.Code = code
	v.Key = key
	v.WindowsVirtualKeyCode = windowsVirtualKeyCode
	v.NativeVirtualKeyCode = nativeVirtualKeyCode
	v.AutoRepeat = autoRepeat
	v.IsKeypad = isKeypad
	v.IsSystemKey = isSystemKey
	v.Location = location
	return c.DispatchKeyEventWithParams(&v)
}

type InputDispatchMouseEventParams struct {
	// Type of the mouse event.
	TheType string `json:"type"`
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y float64 `json:"y"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp float64 `json:"timestamp,omitempty"`
	// Mouse button (default: "none").
	Button string `json:"button,omitempty"`
	// Number of times the mouse button was clicked (default: 0).
	ClickCount int `json:"clickCount,omitempty"`
	// X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaX float64 `json:"deltaX,omitempty"`
	// Y delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY float64 `json:"deltaY,omitempty"`
}

// DispatchMouseEventWithParams - Dispatches a mouse event to the page.
func (c *Input) DispatchMouseEventWithParams(v *InputDispatchMouseEventParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchMouseEvent", Params: v})
}

// DispatchMouseEvent - Dispatches a mouse event to the page.
// type - Type of the mouse event.
// x - X coordinate of the event relative to the main frame's viewport in CSS pixels.
// y - Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred.
// button - Mouse button (default: "none").
// clickCount - Number of times the mouse button was clicked (default: 0).
// deltaX - X delta in CSS pixels for mouse wheel event (default: 0).
// deltaY - Y delta in CSS pixels for mouse wheel event (default: 0).
func (c *Input) DispatchMouseEvent(theType string, x float64, y float64, modifiers int, timestamp float64, button string, clickCount int, deltaX float64, deltaY float64) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchMouseEventParams
	v.TheType = theType
	v.X = x
	v.Y = y
	v.Modifiers = modifiers
	v.Timestamp = timestamp
	v.Button = button
	v.ClickCount = clickCount
	v.DeltaX = deltaX
	v.DeltaY = deltaY
	return c.DispatchMouseEventWithParams(&v)
}

type InputDispatchTouchEventParams struct {
	// Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
	TheType string `json:"type"`
	// Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
	TouchPoints []*InputTouchPoint `json:"touchPoints"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp float64 `json:"timestamp,omitempty"`
}

// DispatchTouchEventWithParams - Dispatches a touch event to the page.
func (c *Input) DispatchTouchEventWithParams(v *InputDispatchTouchEventParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.dispatchTouchEvent", Params: v})
}

// DispatchTouchEvent - Dispatches a touch event to the page.
// type - Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
// touchPoints - Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred.
func (c *Input) DispatchTouchEvent(theType string, touchPoints []*InputTouchPoint, modifiers int, timestamp float64) (*gcdmessage.ChromeResponse, error) {
	var v InputDispatchTouchEventParams
	v.TheType = theType
	v.TouchPoints = touchPoints
	v.Modifiers = modifiers
	v.Timestamp = timestamp
	return c.DispatchTouchEventWithParams(&v)
}

type InputEmulateTouchFromMouseEventParams struct {
	// Type of the mouse event.
	TheType string `json:"type"`
	// X coordinate of the mouse pointer in DIP.
	X int `json:"x"`
	// Y coordinate of the mouse pointer in DIP.
	Y int `json:"y"`
	// Mouse button.
	Button string `json:"button"`
	// Time at which the event occurred (default: current time).
	Timestamp float64 `json:"timestamp,omitempty"`
	// X delta in DIP for mouse wheel event (default: 0).
	DeltaX float64 `json:"deltaX,omitempty"`
	// Y delta in DIP for mouse wheel event (default: 0).
	DeltaY float64 `json:"deltaY,omitempty"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Number of times the mouse button was clicked (default: 0).
	ClickCount int `json:"clickCount,omitempty"`
}

// EmulateTouchFromMouseEventWithParams - Emulates touch event from the mouse event parameters.
func (c *Input) EmulateTouchFromMouseEventWithParams(v *InputEmulateTouchFromMouseEventParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.emulateTouchFromMouseEvent", Params: v})
}

// EmulateTouchFromMouseEvent - Emulates touch event from the mouse event parameters.
// type - Type of the mouse event.
// x - X coordinate of the mouse pointer in DIP.
// y - Y coordinate of the mouse pointer in DIP.
// button - Mouse button.
// timestamp - Time at which the event occurred (default: current time).
// deltaX - X delta in DIP for mouse wheel event (default: 0).
// deltaY - Y delta in DIP for mouse wheel event (default: 0).
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// clickCount - Number of times the mouse button was clicked (default: 0).
func (c *Input) EmulateTouchFromMouseEvent(theType string, x int, y int, button string, timestamp float64, deltaX float64, deltaY float64, modifiers int, clickCount int) (*gcdmessage.ChromeResponse, error) {
	var v InputEmulateTouchFromMouseEventParams
	v.TheType = theType
	v.X = x
	v.Y = y
	v.Button = button
	v.Timestamp = timestamp
	v.DeltaX = deltaX
	v.DeltaY = deltaY
	v.Modifiers = modifiers
	v.ClickCount = clickCount
	return c.EmulateTouchFromMouseEventWithParams(&v)
}

type InputSetIgnoreInputEventsParams struct {
	// Ignores input events processing when set to true.
	Ignore bool `json:"ignore"`
}

// SetIgnoreInputEventsWithParams - Ignores input events (useful while auditing page).
func (c *Input) SetIgnoreInputEventsWithParams(v *InputSetIgnoreInputEventsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.setIgnoreInputEvents", Params: v})
}

// SetIgnoreInputEvents - Ignores input events (useful while auditing page).
// ignore - Ignores input events processing when set to true.
func (c *Input) SetIgnoreInputEvents(ignore bool) (*gcdmessage.ChromeResponse, error) {
	var v InputSetIgnoreInputEventsParams
	v.Ignore = ignore
	return c.SetIgnoreInputEventsWithParams(&v)
}

type InputSynthesizePinchGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`
	// Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	ScaleFactor float64 `json:"scaleFactor"`
	// Relative pointer speed in pixels per second (default: 800).
	RelativeSpeed int `json:"relativeSpeed,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
	GestureSourceType string `json:"gestureSourceType,omitempty"`
}

// SynthesizePinchGestureWithParams - Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
func (c *Input) SynthesizePinchGestureWithParams(v *InputSynthesizePinchGestureParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.synthesizePinchGesture", Params: v})
}

// SynthesizePinchGesture - Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// scaleFactor - Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
// relativeSpeed - Relative pointer speed in pixels per second (default: 800).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
func (c *Input) SynthesizePinchGesture(x float64, y float64, scaleFactor float64, relativeSpeed int, gestureSourceType string) (*gcdmessage.ChromeResponse, error) {
	var v InputSynthesizePinchGestureParams
	v.X = x
	v.Y = y
	v.ScaleFactor = scaleFactor
	v.RelativeSpeed = relativeSpeed
	v.GestureSourceType = gestureSourceType
	return c.SynthesizePinchGestureWithParams(&v)
}

type InputSynthesizeScrollGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`
	// The distance to scroll along the X axis (positive to scroll left).
	XDistance float64 `json:"xDistance,omitempty"`
	// The distance to scroll along the Y axis (positive to scroll up).
	YDistance float64 `json:"yDistance,omitempty"`
	// The number of additional pixels to scroll back along the X axis, in addition to the given distance.
	XOverscroll float64 `json:"xOverscroll,omitempty"`
	// The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
	YOverscroll float64 `json:"yOverscroll,omitempty"`
	// Prevent fling (default: true).
	PreventFling bool `json:"preventFling,omitempty"`
	// Swipe speed in pixels per second (default: 800).
	Speed int `json:"speed,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
	GestureSourceType string `json:"gestureSourceType,omitempty"`
	// The number of times to repeat the gesture (default: 0).
	RepeatCount int `json:"repeatCount,omitempty"`
	// The number of milliseconds delay between each repeat. (default: 250).
	RepeatDelayMs int `json:"repeatDelayMs,omitempty"`
	// The name of the interaction markers to generate, if not empty (default: "").
	InteractionMarkerName string `json:"interactionMarkerName,omitempty"`
}

// SynthesizeScrollGestureWithParams - Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
func (c *Input) SynthesizeScrollGestureWithParams(v *InputSynthesizeScrollGestureParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.synthesizeScrollGesture", Params: v})
}

// SynthesizeScrollGesture - Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// xDistance - The distance to scroll along the X axis (positive to scroll left).
// yDistance - The distance to scroll along the Y axis (positive to scroll up).
// xOverscroll - The number of additional pixels to scroll back along the X axis, in addition to the given distance.
// yOverscroll - The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
// preventFling - Prevent fling (default: true).
// speed - Swipe speed in pixels per second (default: 800).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
// repeatCount - The number of times to repeat the gesture (default: 0).
// repeatDelayMs - The number of milliseconds delay between each repeat. (default: 250).
// interactionMarkerName - The name of the interaction markers to generate, if not empty (default: "").
func (c *Input) SynthesizeScrollGesture(x float64, y float64, xDistance float64, yDistance float64, xOverscroll float64, yOverscroll float64, preventFling bool, speed int, gestureSourceType string, repeatCount int, repeatDelayMs int, interactionMarkerName string) (*gcdmessage.ChromeResponse, error) {
	var v InputSynthesizeScrollGestureParams
	v.X = x
	v.Y = y
	v.XDistance = xDistance
	v.YDistance = yDistance
	v.XOverscroll = xOverscroll
	v.YOverscroll = yOverscroll
	v.PreventFling = preventFling
	v.Speed = speed
	v.GestureSourceType = gestureSourceType
	v.RepeatCount = repeatCount
	v.RepeatDelayMs = repeatDelayMs
	v.InteractionMarkerName = interactionMarkerName
	return c.SynthesizeScrollGestureWithParams(&v)
}

type InputSynthesizeTapGestureParams struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float64 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float64 `json:"y"`
	// Duration between touchdown and touchup events in ms (default: 50).
	Duration int `json:"duration,omitempty"`
	// Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	TapCount int `json:"tapCount,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
	GestureSourceType string `json:"gestureSourceType,omitempty"`
}

// SynthesizeTapGestureWithParams - Synthesizes a tap gesture over a time period by issuing appropriate touch events.
func (c *Input) SynthesizeTapGestureWithParams(v *InputSynthesizeTapGestureParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Input.synthesizeTapGesture", Params: v})
}

// SynthesizeTapGesture - Synthesizes a tap gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// duration - Duration between touchdown and touchup events in ms (default: 50).
// tapCount - Number of times to perform the tap (e.g. 2 for double tap, default: 1).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type). enum values: default, touch, mouse
func (c *Input) SynthesizeTapGesture(x float64, y float64, duration int, tapCount int, gestureSourceType string) (*gcdmessage.ChromeResponse, error) {
	var v InputSynthesizeTapGestureParams
	v.X = x
	v.Y = y
	v.Duration = duration
	v.TapCount = tapCount
	v.GestureSourceType = gestureSourceType
	return c.SynthesizeTapGestureWithParams(&v)
}
