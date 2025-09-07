'use client';

import { PropsWithChildren } from 'react';
import {
	Sidebar,
	SidebarContent,
	SidebarGroup,
	SidebarGroupContent,
	SidebarGroupLabel,
	SidebarMenu,
	SidebarMenuButton,
	SidebarMenuItem,
	SidebarProvider,
} from '@/component/biz/workspace-layout';

const items = [
	{
		title: 'Home',
		url: '#',
	},
	{
		title: 'Inbox',
		url: '#',
	},
	{
		title: 'Calendar',
		url: '#',
	},
	{
		title: 'Search',
		url: '#',
	},
	{
		title: 'Settings',
		url: '#',
	},
];

export default function WorkSpaceLayout({ children }: Readonly<PropsWithChildren>) {
	return (
		<SidebarProvider className='flex h-screen w-full items-center overflow-hidden'>
			<Sidebar>
				<SidebarContent>
					<SidebarGroup>
						<SidebarGroupLabel>Application</SidebarGroupLabel>
						<SidebarGroupContent>
							<SidebarMenu>
								{items.map((item) => (
									<SidebarMenuItem key={item.title}>
										<SidebarMenuButton asChild>
											<a href={item.url}>
												<span>{item.title}</span>
											</a>
										</SidebarMenuButton>
									</SidebarMenuItem>
								))}
							</SidebarMenu>
						</SidebarGroupContent>
					</SidebarGroup>
				</SidebarContent>
			</Sidebar>
			<div className='h-screen min-w-0 flex-1'>{children}</div>
		</SidebarProvider>
	);
}
