package main 


import (
    rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
)

type GameObject struct {
    rec rl.Rectangle;
    col rl.Color;
}

type Game struct {
    objects []GameObject;
    player GameObject;
    width int;
    height int;
    current_time float64;
}

func (self *Game) init_rl() {
    // rl.SetTraceLogLevel(rl.LogError);
    rl.InitWindow(int32(self.width), int32(self.height), "Hello Seaman!");
    rl.SetTargetFPS(60);
    rl.SetExitKey(rl.KeyQ);
}

func (self *Game) close_rl() {
    rl.CloseWindow();
}

func (self *Game) game_loop() {
	for !rl.WindowShouldClose() {
        self.logic_rl();
        self.draw_rl();
	}
}

func (self *Game) draw_rl() {
    rl.BeginDrawing()
    // draw bg
    rl.ClearBackground(rl.GetColor(0x04aaddff))

    // draw player
    rl.DrawRectangleRec(self.player.rec, self.player.col);

    // draw pipe
    for _, obj := range self.objects {
        rl.DrawRectangleRec(obj.rec, obj.col);
    }

    rl.EndDrawing()
}

func (self *Game) logic_rl() {
    // PIPE LOGIC
    time := rl.GetTime();
    if time - self.current_time > 1.5 {
        pipe_wide := float32(rl.GetRandomValue(150, 350));
        top_pipe := GameObject{self.create_pipe(), rl.GetColor(0x00ff00ff)}
        bottom_pipe := GameObject{
            rl.NewRectangle(
                top_pipe.rec.X,
                top_pipe.rec.Height + pipe_wide,
                top_pipe.rec.Width,
                float32(self.height) - top_pipe.rec.Height,
                ),
            rl.GetColor(0x00ff00ff),
        }
        if bottom_pipe.rec.Y > float32(self.height) {
            bottom_pipe.rec.Y = float32(self.height);
        }
        self.objects = append(self.objects, top_pipe);
        self.objects = append(self.objects, bottom_pipe);
        self.current_time = rl.GetTime();
    }
    delta := rl.GetFrameTime();
    const SPEED float32 = 200.0;
    for i := range self.objects {
        self.objects[i].rec.X -= SPEED * delta;
    }
    for i := len(self.objects) - 1; i >= 0; i-- {
        if self.objects[i].rec.X < -100 {
            self.objects = slices.Delete(self.objects, i, i+1);
        }
    }

    // PLAYER LOGIC
    if rl.IsKeyDown(rl.KeySpace) {
        self.player.rec.Y -= (SPEED + 100) * delta;
    } else {
        self.player.rec.Y += (SPEED + SPEED * 0.3) * delta;
    }

    // CHECK COLLISION
    for _, obj := range self.objects {
        if self.player.rec.X > obj.rec.X &&
            obj.rec.X < (float32(self.width)/2) {
            continue;
        }
        if rl.CheckCollisionRecs(self.player.rec, obj.rec) {
            rl.CloseWindow();
        }
    }
    return;
}

func (self *Game) create_pipe() rl.Rectangle {
    return rl.NewRectangle(float32(self.width) + 100, 0, 75, float32(rl.GetRandomValue(200, 500)));
}

func (self *Game) init_game_objects() {
    self.player = GameObject{rl.NewRectangle(20, 0, 60, 60), rl.GetColor(0xffff00ff)};
    self.current_time = rl.GetTime();
    return;
}
