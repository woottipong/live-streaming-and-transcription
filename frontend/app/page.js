import {
  ArrowUpRightIcon,
  CaptionsIcon,
  Mic2Icon,
  PanelTopIcon,
  RadioTowerIcon,
} from "lucide-react";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

const highlights = [
  {
    title: "Admin Control",
    description: "Create sessions, prepare rooms, and manage source metadata.",
    icon: PanelTopIcon,
  },
  {
    title: "Live Publishing",
    description: "Support browser publishers today and RTMP ingress in the next step.",
    icon: RadioTowerIcon,
  },
  {
    title: "Realtime Captions",
    description: "Stream audio into the transcription worker and fan out captions live.",
    icon: CaptionsIcon,
  },
];

export default function HomePage() {
  return (
    <main className="min-h-screen px-6 py-8 md:px-10 md:py-10">
      <div className="mx-auto flex max-w-6xl flex-col gap-6">
        <Card className="border-none bg-transparent py-0 shadow-none ring-0">
          <CardHeader className="gap-4 px-0">
            <div className="flex flex-wrap items-center gap-3">
              <Badge variant="secondary">Frontend ready</Badge>
              <Badge variant="outline">shadcn/ui</Badge>
            </div>
            <CardTitle className="max-w-3xl text-4xl leading-none tracking-tight md:text-6xl">
              Realtime streaming UI foundation for admin, publisher, and viewer
              flows.
            </CardTitle>
            <CardDescription className="max-w-2xl text-base leading-7">
              The frontend now uses shadcn components as the shared UI layer, so
              the next tasks can build on consistent cards, actions, and status
              patterns instead of custom markup.
            </CardDescription>
          </CardHeader>
          <CardFooter className="flex flex-wrap gap-3 border-none bg-transparent px-0 pt-0">
            <Button size="lg">
              Explore session setup
              <ArrowUpRightIcon data-icon="inline-end" />
            </Button>
            <Button variant="outline" size="lg">
              Worker integration next
              <Mic2Icon data-icon="inline-end" />
            </Button>
          </CardFooter>
        </Card>

        <section className="grid gap-4 lg:grid-cols-[minmax(0,1.4fr)_minmax(0,1fr)]">
          <Card className="border-border/60 bg-card/90 backdrop-blur">
            <CardHeader>
              <CardTitle>MVP delivery path</CardTitle>
              <CardDescription>
                The first frontend slice stays intentionally narrow so we can
                connect LiveKit and the transcription worker without redesigning
                the UI surface later.
              </CardDescription>
              <CardAction>
                <Badge>Phase 1</Badge>
              </CardAction>
            </CardHeader>
            <CardContent>
              <div className="grid gap-3 md:grid-cols-3">
                {highlights.map(({ title, description, icon: Icon }) => (
                  <Card key={title} size="sm" className="bg-background/80">
                    <CardHeader>
                      <CardTitle className="flex items-center gap-2 text-sm">
                        <Icon />
                        {title}
                      </CardTitle>
                      <CardDescription>{description}</CardDescription>
                    </CardHeader>
                  </Card>
                ))}
              </div>
            </CardContent>
            <CardFooter className="justify-between gap-3">
              <p className="text-muted-foreground">
                Session CRUD and LiveKit wiring are the next focused tasks.
              </p>
              <Badge variant="outline">Monorepo scaffold done</Badge>
            </CardFooter>
          </Card>

          <Card className="bg-card/95">
            <CardHeader>
              <CardTitle>Current setup snapshot</CardTitle>
              <CardDescription>
                Verified services and the remaining blocker for closing the
                scaffold task.
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex flex-col gap-3">
                <div className="flex items-center justify-between gap-3 rounded-lg border bg-background/70 px-3 py-2">
                  <span>Backend service</span>
                  <Badge>Port 8080</Badge>
                </div>
                <div className="flex items-center justify-between gap-3 rounded-lg border bg-background/70 px-3 py-2">
                  <span>Worker service</span>
                  <Badge>Port 8081</Badge>
                </div>
                <div className="flex items-center justify-between gap-3 rounded-lg border bg-background/70 px-3 py-2">
                  <span>Frontend dev server</span>
                  <Badge>Port 3000</Badge>
                </div>
                <div className="flex items-center justify-between gap-3 rounded-lg border bg-background/70 px-3 py-2">
                  <span>Docker verification</span>
                  <Badge variant="outline">Daemon pending</Badge>
                </div>
              </div>
            </CardContent>
            <CardFooter>
              <Button variant="secondary" className="w-full">
                Ready for admin page task
              </Button>
            </CardFooter>
          </Card>
        </section>
      </div>
    </main>
  );
}
